import { useState, useEffect } from 'react'
import './App.css'
import './DashPlayer'
import DashPlayer from './DashPlayer'
import Playlist from './Playlist'
import ArtistComponent from './Artist'
import { Client } from './client'
import type { Artist, Album } from './client'

function App() {
	const [currentSongUrl, setCurrentSongUrl] = useState<string | null>(null);
	const [currentArtist, setCurrentArtist] = useState<string | null>(null)
	const [artists, setArtists] = useState<Artist[]>([])
	const [albums, setAlbums] = useState<Album[]>([])

	const client = new Client()

	const onSongEnd = () => {
		if (currentSongUrl == null) {
			return
		}
		for (let i = 0; i < albums.length; i++) {
			const album = albums[i];
			for (let j = 0; j < album.songs.length; j++) {
				const song = album.songs[j];
				if (song.file == currentSongUrl) {
					const order = j + 1
					if (order >= album.songs.length) {
						setCurrentSongUrl(album.songs[0].file)
						return
					}
					setCurrentSongUrl(album.songs[order].file)
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
		const nextSong = getNextSong(albums, currentSongUrl)
		if (nextSong == null) {
			return
		}
		setCurrentSongUrl(nextSong)
	}

	const onPreviousSongButtonClick = () => {
		if (currentSongUrl == null) {
			return
		}
		const previousSong = getPreviousSong(albums, currentSongUrl)
		if (previousSong == null) {
			return
		}
		setCurrentSongUrl(previousSong)
	}

	useEffect(() => {
		const fetchAlbums = async () => {
			try {
				if (currentArtist) {
					const albs = await client.getAlbumsByArtist(currentArtist)
					setAlbums(albs)
				}
			} catch (err) {
				console.error("error fetching albums by artist: ", err)
			}

		}
		fetchAlbums()
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
			<nav className='navbar'>
				<h1>Navbar</h1>
			</nav>
			<main className='content'>
				<div className='list-container rad-shadow'>
					<ArtistComponent artists={artists} setCurrentArtist={setCurrentArtist} currentArtist={currentArtist} />
				</div>
				<div className='list-container rad-shadow'>
					{albums.length > 0 ? (
						albums.map((album) => (
							<Playlist album={album} setCurrentSongUrl={setCurrentSongUrl} highLightedSong={currentSongUrl} />
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

function getPreviousSong(albums: Album[], currentSongUrl: string): string | null {
	for (let i = 0; i < albums.length; i++) {
		const album = albums[i];
		for (let j = 0; j < album.songs.length; j++) {
			const song = album.songs[j];
			if (song.file == currentSongUrl) {
				const order = j - 1
				if (order < 0) {
					return album.songs[album.songs.length - 1].file
				}
				return album.songs[order].file
			}
		}
	}
	return null
}
function getNextSong(albums: Album[], currentSongUrl: string): string | null {
	for (let i = 0; i < albums.length; i++) {
		const album = albums[i];
		for (let j = 0; j < album.songs.length; j++) {
			const song = album.songs[j];
			if (song.file == currentSongUrl) {
				const order = j + 1
				if (order >= album.songs.length) {
					return album.songs[0].file
				}
				return album.songs[order].file
			}
		}
	}
	return null
}

export default App
