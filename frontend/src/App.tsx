import { useState, useEffect } from 'react'
import './App.css'
import './DashPlayer'
import DashPlayer from './DashPlayer'
import Playlist from './Playlist'
import ArtistComponent from './Artist'
import { Client } from './client'
import type { Artist, Release } from './client'

function App() {
	const [currentTrackUrl, setCurrentTrackUrl] = useState<string | null>(null);
	const [currentArtist, setCurrentArtist] = useState<string | null>(null)
	const [artists, setArtists] = useState<Artist[]>([])
	const [releases, setReleases] = useState<Release[]>([])

	const client = new Client()

	/*
	const onTrackEnd = () => {
		if (currentTrackUrl == null) {
			return
		}
		for (let i = 0; i < releases.length; i++) {
			const release = releases[i];
			for (let j = 0; j < release.tracks.length; j++) {
				const track = release.tracks[j];
				if (track.file == currentTrackUrl) {
					const order = j + 1
					if (order >= release.tracks.length) {
						setCurrentTrackUrl(release.tracks[0].file)
						return
					}
					setCurrentTrackUrl(release.tracks[order].file)
					return
				}

			}
		}
		return
	}
	*/
	const onNextTrackButtonClick = () => {
		if (currentTrackUrl == null) {
			return
		}
		const nextTrack = getNextTrack(releases, currentTrackUrl)
		if (nextTrack == null) {
			return
		}
		setCurrentTrackUrl(nextTrack)
	}

	const onPreviousTrackButtonClick = () => {
		if (currentTrackUrl == null) {
			return
		}
		const previousTrack = getPreviousTrack(releases, currentTrackUrl)
		if (previousTrack == null) {
			return
		}
		setCurrentTrackUrl(previousTrack)
	}

	useEffect(() => {
		const fetchReleases = async () => {
			try {
				if (currentArtist) {
					const albs = await client.getReleasesByArtist(currentArtist)
					setReleases(albs)
				}
			} catch (err) {
				console.error("error fetching releases by artist: ", err)
			}

		}
		fetchReleases()
	}, [currentArtist])

	useEffect(() => {
		const fetchArtists = async () => {
			try {
				const ars = await client.getAllArtists()
				setArtists(ars)

			} catch (err) {
				console.error("error fetching all artists: ", err)
			}
		}
		fetchArtists()
	}, [])


	return (
		<div className='app-container'>
			{/*
			<nav className='navbar'>
				<h1>Navbar</h1>
			</nav>
			*/}
			<main className='content'>
				<div className='list-container rad-shadow'>
					<ArtistComponent artists={artists} setCurrentArtist={setCurrentArtist} currentArtist={currentArtist} />
				</div>
				<div className='list-container rad-shadow'>
					{releases.length > 0 ? (
						releases.map((release) => (
							<Playlist release={release} setCurrentTrackUrl={setCurrentTrackUrl} highLightedTrack={currentTrackUrl} />
						))
					) : (
						<strong>Pick a artist</strong>
					)}
				</div>
			</main>
			<footer className='footer'>
				{currentTrackUrl ? (
					<DashPlayer
						src={currentTrackUrl}
						getNextSrc={() => getNextTrack(releases, currentTrackUrl)}
						onNextTrack={onNextTrackButtonClick}
						onPreviousTrack={onPreviousTrackButtonClick}
						autoplay
					/>
				) : (
					<DashPlayer
						src=''
						getNextSrc={() => null}
						onNextTrack={onNextTrackButtonClick}
						onPreviousTrack={onPreviousTrackButtonClick}
						autoplay={false}
					/>
				)}
			</footer>
		</div >
	)
}

function getPreviousTrack(releases: Release[], currentTrackUrl: string): string | null {
	for (let i = 0; i < releases.length; i++) {
		const release = releases[i];
		for (let j = 0; j < release.tracks.length; j++) {
			const track = release.tracks[j];
			if (track.file == currentTrackUrl) {
				const order = j - 1
				if (order < 0) {
					return release.tracks[release.tracks.length - 1].file
				}
				return release.tracks[order].file
			}
		}
	}
	return null
}
function getNextTrack(releases: Release[], currentTrackUrl: string): string | null {
	for (let i = 0; i < releases.length; i++) {
		const release = releases[i];
		for (let j = 0; j < release.tracks.length; j++) {
			const track = release.tracks[j];
			if (track.file == currentTrackUrl) {
				const order = j + 1
				if (order >= release.tracks.length) {
					return release.tracks[0].file
				}
				return release.tracks[order].file
			}
		}
	}
	return null
}

export default App
