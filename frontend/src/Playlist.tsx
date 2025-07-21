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
					<h2 style={{ textAlign: 'center', color: 'var(--text1)' }}>{album?.name}</h2>
					<table style={{ width: '100%', borderCollapse: 'collapse' }}>
						<thead style={{ borderBottom: '2px solid var(--surface2)' }}>
							<tr style={{ color: 'var(--text2)' }}>
								<th style={{ textAlign: 'left' }}>Song</th>
								<th style={{ textAlign: 'right' }}>Duration</th>
							</tr>
						</thead>
						<tbody>
							{album.songs.map(song => (
								<tr onClick={() => setCurrentSongUrl(song.file)} style={{ cursor: 'pointer' }}>
									<td key={song.file} style={{ textAlign: 'left', color: 'var(--text1)' }}>
										{highLightedSong != '' && song.file == highLightedSong ? (
											<strong className='current-song'>{song.name}</strong>
										) : (
											<strong>{song.name}</strong>
										)}
									</td>
									<td style={{ textAlign: 'right', color: 'var(--text2)' }}>{formatTime(song.duration)}</td>
								</tr>
							))}
						</tbody>
					</table>
				</div>
			)}
		</div>
	);
};

function formatTime(time: number) {
	const minutes = Math.floor(time / 60);
	const seconds = Math.floor(time % 60).toString().padStart(2, '0');
	return `${minutes}:${seconds}`;
}


export default Playlist;

