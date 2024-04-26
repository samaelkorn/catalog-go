import { urlNext } from '@/utils/url'
export async function getProducts() {
    const res = await fetch(urlNext('api', 'products'))

    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}