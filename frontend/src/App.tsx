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
				<div className='list-container'>
					<ArtistComponent artists={artists} setCurrentArtist={setCurrentArtist} />
				</div>
				<div className='list-container'>
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
					<DashPlayer src={currentSongUrl} onSongEnd={onSongEnd} autoplay />
				) : (
					<DashPlayer src='' onSongEnd={onSongEnd}/>
				)}
			</footer>
		</div>
	)
}

export default App
