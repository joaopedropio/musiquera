import React, { useEffect, useState } from 'react';
import type { Album } from './client';
import { Client } from './client'

interface PlaylistProps {
	setCurrentSongUrl: (url: string) => void;
}
const Playlist: React.FC<PlaylistProps> = ({ setCurrentSongUrl }) => {
	const [album, setAlbum] = useState<Album | null>(null);
	const [loading, setLoading] = useState<boolean>(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		const client = new Client()
		const fetchAlbum = async () => {
			try {
				const a = await client.getMostRecentAlbum()
				setAlbum(a)
				if (a.songs && a.songs.length > 0) {
					console.log("called set current song url")
					setCurrentSongUrl(a.songs[0].file)
				}
			} catch (err) {
				setError((err as Error).message);
			} finally {
				setLoading(false)
			}
		}
		fetchAlbum();
	}, [setCurrentSongUrl]);

	if (loading) return <div>Loading album...</div>;
	if (error) return <div>Error: {error}</div>;
	console.log('album:', album);
	console.log('songs:', album?.songs);
	console.log('songs length:', album?.songs?.length);

	return (
		<div>
			{!album || !album.songs || album.songs.length === 0 ? (
				<p>No songs found.</p>
			) : (
				<div>
					<h1>{album?.artist}</h1>
					<h2>{album?.name}</h2>
					<ul>
						{album.songs.map(song => (
							<li key={song.file} onClick={() => setCurrentSongUrl(song.file)} style={{ cursor: 'pointer' }}>
								<strong>{song.name}</strong>
							</li>
						))}
					</ul>
				</div>
			)}
		</div>
	);
};

export default Playlist;

