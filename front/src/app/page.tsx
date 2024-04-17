import Link from 'next/link'

export default function Home() {
  return (
    <div className="p-9">
      <Link href="/catalog" className="shadow-md mx-auto w-1/4 flex h-32 text-2xl text-white bg-gray justify-center items-center">
        <div>
          Каталог
        </div>
      </Link>
    </div>
  )
}
