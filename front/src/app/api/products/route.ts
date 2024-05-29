import { urlApi } from '@/utils/url'
export async function GET(request: Request) {
  const { searchParams } = new URL(request.url)
  const params = searchParams && searchParams.size ? '?' + searchParams.toString() : ''
  const res = await fetch(urlApi('products' + params), {
    headers: {
      'Content-Type': 'application/json'
    },
  })
  const data = await res.json()

  return Response.json(data)
}