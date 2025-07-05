import axios from 'axios'

export type Song = {
	name: string;
	file: string;
}

export type Album = {
	name: string;
	artist: string;
	releaseDate: string;
	songs: Song[];
}

export type Artist = {
	name: string;
}

export class Client {
	constructor() {

	}

	async getAllArtists(): Promise<Artist[]> {
		const response = await axios.get<Artist[]>('/api/artist/')
		return response.data
	}

	async getAlbumsByArtist(artistName: string): Promise<Album[]> {
		const response = await axios.get<Album[]>('/api/album/byArtist/' + artistName)
		return response.data
	}

	async getMostRecentAlbum(): Promise<Album> {
		const resp = await axios.get<Album>('/api/album/mostRecent')
		return resp.data
	}
}
