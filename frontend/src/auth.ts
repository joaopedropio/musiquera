export async function isAuthenticated(): Promise<boolean> {
	try {
		const res = await fetch('/auth/check', {
			credentials: 'include',
		})
		return res.ok
	} catch {
		return false
	}
}

