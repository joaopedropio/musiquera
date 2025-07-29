// auth.ts
export async function isAuthenticated(): Promise<boolean> {
	try {
		const res = await fetch('/auth/check', {
			credentials: 'include', // needed to send cookie
		})
		if (!res.ok) return false

		const data = await res.json()
		return data.authenticated === true
	} catch {
		return false
	}
}

