'use client'
import { useState, useEffect } from 'react'
import { useSearchParams } from 'next/navigation'
import Link from 'next/link'
import Image from 'next/image'
import { getProducts } from '@/api'

interface TP {
    data: TProduct[]
    search: string
}

export default function Catalog({ data, search }: TP) {
    const searchParams = useSearchParams()
    const [list, setList] = useState<TProduct[]>(data)
    const [stateSearch, setStateSearch] = useState(search)

    useEffect(() => {
        if (stateSearch !== searchParams.toString()) {
            setStateSearch(searchParams.toString())
            getProducts(searchParams)
                .then(({ data }) => setList(data))
        }
    }, [stateSearch, searchParams])

    return (
        <div className="grid grid-cols-1 md:grid-cols-4 gap-4 mt-5">
            {list.map((item) =>
                <div key={item.id} className="rounded-lg shadow-sm overflow-hidden">
                    <div>
                        <Image src={item.image} alt="" width={600} height={400} />
                    </div>
                    <div className="px-3 py-4">
                        <div className="text-2xl">
                            {item.name}
                        </div>
                        <div className="text-xl">
                            {item.price} €
                        </div>
                        <div className="mt-2 text-slate-500">
                            <span>{item.color.name}</span>
                        </div>
                        <div className="mt-2">
                            <span className="rounded-lg bg-green-50 text-green-900 p-1">{item.status.name}</span>
                        </div>
                        <div className="mt-6">
                            <Link href={`/catalog/${item.id}`} className="rounded-lg bg-pink-900 text-white py-3 px-2 w-full block">
                                Подбробнее
                            </Link>
                        </div>
                    </div>
                </div>
            )}
        </div>
    )
}