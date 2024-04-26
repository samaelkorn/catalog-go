import { urlNext } from '@/utils/url'
export async function getColors() {
    const res = await fetch(urlNext('api', 'colors'))

    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export async function getStatuses() {
    const res = await fetch(urlNext('api', 'statuses'))

    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}