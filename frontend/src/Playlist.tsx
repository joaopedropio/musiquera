import type { Album } from './client';
import './Playlist.css'

interface PlaylistProps {
	setCurrentSongUrl: (url: string) => void;
	highLightedSong: string | null;
	album: Album;
}
const Playlist: React.FC<PlaylistProps> = ({ album, setCurrentSongUrl, highLightedSong }) => {
	return (
		<div className='playlist'>
			{!album || !album.songs || album.songs.length === 0 ? (
				<p>No songs found.</p>
			) : (
				<div>
					<h2>{album?.name}</h2>
					<ul>
						{album.songs.map(song => (
							<li key={song.file} onClick={() => setCurrentSongUrl(song.file)} style={{ cursor: 'pointer' }}>
								{highLightedSong != '' && song.file == highLightedSong ? (
									<strong style={{ color: 'green' }}>{song.name}</strong>
								) : (
									<strong>{song.name}</strong>
								)}
							</li>
						))}
					</ul>
				</div>
			)}
		</div>
	);
};

export default Playlist;

