import { useState, useEffect } from 'react'
import './App.css'
import './DashPlayer'
import DashPlayer from './DashPlayer'
import Playlist from './Playlist'
import ArtistComponent from './Artist'
import { Client } from './client'
import type { Artist, Release } from './client'

function App() {
	const [currentSongUrl, setCurrentSongUrl] = useState<string | null>(null);
	const [currentArtist, setCurrentArtist] = useState<string | null>(null)
	const [artists, setArtists] = useState<Artist[]>([])
	const [releases, setReleases] = useState<Release[]>([])

	const client = new Client()

	const onSongEnd = () => {
		if (currentSongUrl == null) {
			return
		}
		for (let i = 0; i < releases.length; i++) {
			const release = releases[i];
			for (let j = 0; j < release.songs.length; j++) {
				const song = release.songs[j];
				if (song.file == currentSongUrl) {
					const order = j + 1
					if (order >= release.songs.length) {
						setCurrentSongUrl(release.songs[0].file)
						return
					}
					setCurrentSongUrl(release.songs[order].file)
					return
				}

			}
		}
		return
	}

	const onNextSongButtonClick = () => {
		if (currentSongUrl == null) {
			return
		}
		const nextSong = getNextSong(releases, currentSongUrl)
		if (nextSong == null) {
			return
		}
		setCurrentSongUrl(nextSong)
	}

	const onPreviousSongButtonClick = () => {
		if (currentSongUrl == null) {
			return
		}
		const previousSong = getPreviousSong(releases, currentSongUrl)
		if (previousSong == null) {
			return
		}
		setCurrentSongUrl(previousSong)
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
							<Playlist release={release} setCurrentSongUrl={setCurrentSongUrl} highLightedSong={currentSongUrl} />
						))
					) : (
						<strong>Pick a artist</strong>
					)}
				</div>
			</main>
			<footer className='footer'>
				{currentSongUrl ? (
					<DashPlayer
						src={currentSongUrl}
						onSongEnd={onSongEnd}
						onNextSong={onNextSongButtonClick}
						onPreviousSong={onPreviousSongButtonClick}
						autoplay
					/>
				) : (
					<DashPlayer
						src=''
						onSongEnd={onSongEnd}
						onPreviousSong={onPreviousSongButtonClick}
						onNextSong={onNextSongButtonClick}
					/>
				)}
			</footer>
		</div>
	)
}

function getPreviousSong(releases: Release[], currentSongUrl: string): string | null {
	for (let i = 0; i < releases.length; i++) {
		const release = releases[i];
		for (let j = 0; j < release.songs.length; j++) {
			const song = release.songs[j];
			if (song.file == currentSongUrl) {
				const order = j - 1
				if (order < 0) {
					return release.songs[release.songs.length - 1].file
				}
				return release.songs[order].file
			}
		}
	}
	return null
}
function getNextSong(releases: Release[], currentSongUrl: string): string | null {
	for (let i = 0; i < releases.length; i++) {
		const release = releases[i];
		for (let j = 0; j < release.songs.length; j++) {
			const song = release.songs[j];
			if (song.file == currentSongUrl) {
				const order = j + 1
				if (order >= release.songs.length) {
					return release.songs[0].file
				}
				return release.songs[order].file
			}
		}
	}
	return null
}

export default App
