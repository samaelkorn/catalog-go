import { urlNext } from '@/utils/url'
export async function getProducts(data: URLSearchParams | null = null) {
    const params = data && data.size ? '?' + data.toString() : ''
    const res = await fetch(urlNext('api', 'products') + params)

    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}