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
		<div className='container'>
			<div className='content'>
				<div>
					<ArtistComponent artists={artists} setCurrentArtist={setCurrentArtist} />
				</div>
				<div>
					{albums.length > 0 ? (
						albums.map((album) => (

							<Playlist album={album} setCurrentSongUrl={setCurrentSongUrl} />
						))
					) : (
						<strong>Pick a artist</strong>
					)}
				</div>
			</div>
			<div className='footer'>
				{ currentSongUrl ? (
					<DashPlayer src={currentSongUrl} autoplay />
				):(
					<DashPlayer src='' />
				)}
			</div>
		</div>
	)
}

export default App
