import type { Artist } from './client'
import './Artist.css'

interface Props {
	artists: Artist[]
	setCurrentArtist: (artist: string) => void
}

const ArtistComponent: React.FC<Props> = ({ artists, setCurrentArtist }) => {
	return (
		<div className='artistComponent'>
			<h2 style={{color: 'var(--text1)', textAlign: 'center'}}>Artists</h2>
			{artists.length === 0 ? (
				<p> No artists found. </p>
			) : (
				<ul className='no-dots'>
					{artists.map(artist => (
						<li key={artist.name} onClick={() => setCurrentArtist(artist.name)} style={{ cursor: 'pointer' }}>
							<strong style={{color: 'var(--text2)'}}>{artist.name}</strong>
						</li>
					))}
				</ul>
			)}

		</div>
	)
}

export default ArtistComponent
