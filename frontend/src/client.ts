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

export class Client {
	constructor() {

	}

	async getMostRecentAlbum(): Promise<Album> {
		const response = await axios.get<Album>('/api/album/mostRecent')
		return response.data
	}
}
