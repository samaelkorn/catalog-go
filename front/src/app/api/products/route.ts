import { urlApi } from '@/utils/url'
export async function GET() {
  const res = await fetch(urlApi('products'), {
    headers: {
      'Content-Type': 'application/json'
    },
  })
  const data = await res.json()

  return Response.json(data)
}