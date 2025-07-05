import React, { useEffect, useState } from 'react'
import { Client } from './client'
import type { Artist } from './client'

interface Props {

}

const ArtistComponent: React.FC<Props> = () => {
	const [allArtists, setAllArtists] = useState<Artist[]>([])
	const [loading, setLoading] = useState<boolean>(true)
	const [error, setError] = useState<string | null>(null)

	useEffect(() => {
		const fetchArtists = async () => {
			try {
				const client = new Client()
				const artists = await client.getAllArtists()
				setAllArtists(artists)

			} catch (err) {
				setError((err as Error).message)
			} finally {
				setLoading(false)
			}

		}
		fetchArtists()
	}, [])

	if (loading) return <div>Loading Artist...</div>
	if (error) return <div>Error: {error}</div>

	return (
		<div>
			{allArtists.length === 0 ? (
				<p> No artists found. </p>
			) : (
				<ul>
					{allArtists.map(artist => (
						<li key={artist.name}>
							<strong>{artist.name}</strong>
						</li>
					))}
				</ul>
			)}

		</div>
	)
}

export default ArtistComponent
