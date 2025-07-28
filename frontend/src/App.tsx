import { useState, useEffect } from 'react'
import './App.css'
import './DashPlayer'
import DashPlayer from './DashPlayer'
import DashPlayerLiveSet from './DashPlayerLiveSet'
import Album from './Album'
import ArtistComponent from './Artist'
import { Client } from './client'
import type { Segment, Artist, Release } from './client'
import LiveSet from './LiveSet'

function App() {
	const [currentTrackUrl, setCurrentTrackUrl] = useState<string | null>(null);
	const [currentArtist, setCurrentArtist] = useState<string | null>(null)
	const [currentSegment, setCurrentSegment] = useState<Segment | null>(null)
	const [currentReleaseType, setCurrentReleaseType] = useState<string | null>(null)
	const [artists, setArtists] = useState<Artist[]>([])
	const [releases, setReleases] = useState<Release[]>([])

	const client = new Client()

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

	const onNextSegmentButtonClick = () => {
		if (currentTrackUrl == null) {
			return
		}
		if (currentSegment == null) {
			return
		}
		const nextSegment = getNextSegment(releases, currentTrackUrl, currentSegment)
		if (nextSegment == null) {
			return
		}
		setCurrentSegment(nextSegment)
	}

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
	const onPreviousSegmentButtonClick = () => {
		if (currentTrackUrl == null) {
			return
		}
		if (currentSegment == null) {
			return
		}
	 	const previousSegment = getPreviousSegment(releases, currentTrackUrl, currentSegment)
		if (previousSegment == null) {
			return
		}
		setCurrentSegment(previousSegment)
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
			<main className='content'>
				<div className='list-container rad-shadow'>
					<ArtistComponent artists={artists} setCurrentArtist={setCurrentArtist} currentArtist={currentArtist} />
				</div>
				<div className='list-container rad-shadow'>
					{releases.length > 0 ? (
						releases.map((release) => 
							release.type === 'album' ? (
							<Album release={release} setCurrentTrackUrl={setCurrentTrackUrl} highLightedTrack={currentTrackUrl} setCurrentReleaseType={setCurrentReleaseType}/>
						): (
							<LiveSet release={release} setCurrentTrackUrl={setCurrentTrackUrl} setCurrentSegment={setCurrentSegment} highLightedTrack={currentTrackUrl} setCurrentReleaseType={setCurrentReleaseType}/>
						)
					)) : (
						<strong>Pick a artist</strong>
					)}
				</div>
			</main>
			<footer className='footer'>
				{currentTrackUrl && currentReleaseType === 'album' ? (
					<DashPlayer
						src={currentTrackUrl}
						onTrackEnd={onTrackEnd}
						onNextTrack={onNextTrackButtonClick}
						onPreviousTrack={onPreviousTrackButtonClick}
						autoplay
					/>
				) : currentTrackUrl && currentSegment && currentReleaseType === 'liveSet' ?(
					<DashPlayerLiveSet
						src={currentTrackUrl}
						currentSegment={currentSegment}
						onPreviousSegment={onPreviousSegmentButtonClick}
						onNextSegment={onNextSegmentButtonClick}
					/>
				) : (
					<DashPlayer
						src=''
						onTrackEnd={onTrackEnd}
						onPreviousTrack={onPreviousTrackButtonClick}
						onNextTrack={onNextTrackButtonClick}
					/>
				)}
			</footer>
		</div>
	)
}

function getNextSegment(releases: Release[], currentTrackUrl: string, currentSegment: Segment): Segment| null {
	for (let i = 0; i < releases.length; i++) {
		const release = releases[i];
		for (let j = 0; j < release.tracks.length; j++) {
			const track = release.tracks[j];
			if (track.file == currentTrackUrl) {
				for (let k = 0; k < track.segments.length; k++) {
					if (track.segments[k].name == currentSegment.name) {
						const order = k + 1
						if (order >= track.segments.length) {
							return track.segments[0]
						}
						return track.segments[order]
					}
				}
			}
		}
	}
	return null
}

function getPreviousSegment(releases: Release[], currentTrackUrl: string, currentSegment: Segment): Segment| null {
	for (let i = 0; i < releases.length; i++) {
		const release = releases[i];
		for (let j = 0; j < release.tracks.length; j++) {
			const track = release.tracks[j];
			if (track.file == currentTrackUrl) {
				for (let k = 0; k < track.segments.length; k++) {
					if (track.segments[k].name == currentSegment.name) {
						const order = k - 1
						if (order < 0) {
							return track.segments[track.segments.length - 1]
						}
						return track.segments[order]
					}
				}
			}
		}
	}
	return null
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
