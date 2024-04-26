'use client'
import { useSearchParams, usePathname, useRouter } from 'next/navigation'
import { useCallback } from 'react'
interface TP {
    statuses: TStatus[]
    colors: TColor[]
}

export default function Filter({ statuses, colors }: TP) {
    const router = useRouter()
    const pathname = usePathname()
    const searchParams = useSearchParams()

    const createQueryString = useCallback(
        (name: string, value: string) => {
            const params = new URLSearchParams(searchParams.toString())
            value && params.set(name, value)
            !value && params.delete(name)
            return params.toString()
        },
        [searchParams]
    )

    return (
        <div className="shadow-lg rounded-[10px] py-10 px-5 bg-white">
            <div className="flex">
                <div className="mr-2">
                    <select defaultValue="" className="border border-stone-400 text-sm rounded-[7px] p-2.5 pr-4" onChange={({ target }) => router.push(pathname + '?' + createQueryString('color', target.value))}>
                        <option selected value="">Выберите цвет</option>
                        {colors.map(item =>
                            <option key={item.id} value={item.id}>{item.name}</option>
                        )}
                    </select>
                </div>
                <div>
                    <select defaultValue="" className="border border-stone-400 text-sm rounded-[7px] p-2.5 pr-4" onChange={({ target }) => router.push(pathname + '?' + createQueryString('status', target.value))}>
                        <option selected value="">Выберите статус</option>
                        {statuses.map(item =>
                            <option key={item.id} value={item.id}>{item.name}</option>
                        )}
                    </select>
                </div>
            </div>
        </div>
    )
}