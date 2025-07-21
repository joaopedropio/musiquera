import { useEffect } from 'react'
import type { Artist } from './client'
import './Artist.css'

interface Props {
	artists: Artist[]
	setCurrentArtist: (artist: string) => void
	currentArtist: string | null
}

const ArtistComponent: React.FC<Props> = ({ artists, setCurrentArtist, currentArtist }) => {
	useEffect(() => {
		if (artists.length > 0) {
			setCurrentArtist(artists[0].name)
		}

	}, [artists])
	return (
		<div className='artistComponent'>
			<h2 style={{ color: 'var(--text1)', textAlign: 'center' }}>Artists</h2>
			{artists.length === 0 ? (
				<p> No artists found. </p>
			) : (
				<ul className='no-dots'>
					{artists.map(artist => (
						<li key={artist.name} onClick={() => setCurrentArtist(artist.name)} style={{ cursor: 'pointer' }}>
							{currentArtist != null && currentArtist == artist.name ? (
								<strong style={{ color: 'var(--brand)' }}>{artist.name}</strong>
							) : (
								<strong style={{ color: 'var(--text1)' }}>{artist.name}</strong>
							)}
						</li>
					))}
				</ul>
			)}

		</div>
	)
}

export default ArtistComponent
