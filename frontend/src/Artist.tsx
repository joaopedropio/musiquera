import type { Artist } from './client'
import './Artist.css'

interface Props {
	artists: Artist[]
	setCurrentArtist: (artist: string) => void
}

const ArtistComponent: React.FC<Props> = ({ artists, setCurrentArtist }) => {
	return (
		<div className='artistComponent'>
			{artists.length === 0 ? (
				<p> No artists found. </p>
			) : (
				<ul>
					{artists.map(artist => (
						<li key={artist.name} onClick={() => setCurrentArtist(artist.name)} style={{ cursor: 'pointer' }}>
							<strong>{artist.name}</strong>
						</li>
					))}
				</ul>
			)}

		</div>
	)
}

export default ArtistComponent
