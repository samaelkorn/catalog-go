import Link from 'next/link'

const listNav = [
    {
        title: 'Новые мотоциклы',
        href: '/catalog',
        id: 1
    },
    {
        title: 'Новые мотоциклы',
        href: '/catalog',
        id: 2
    },
    {
        title: 'Новые мотоциклы',
        href: '/catalog',
        id: 3
    }
]

export const Header = () => {
    return (
        <div className="bg-gray h-32">
            <div className="container mx-auto flex flex-col h-full divide-y divide-white text-white divide-slate-400/25">
                <div className="h-1/2 flex justify-between items-center">
                    <div>Москва</div>
                    <div>Мы в социальных сетях</div>
                    <div>Заказать звонок</div>
                </div>
                <div className="h-1/2 flex justify-between items-center">
                    <div>
                        <Link href="/" className="uppercase">Samael-Moto</Link>
                    </div>
                    <nav>
                        {
                            listNav.map((item, index) =>
                                <Link href={item.href} className={listNav.length !== index + 1 ? 'mr-4' : ''} key={item.id}>
                                    {item.title}
                                </Link>
                            )
                        }
                    </nav>
                    <div>Иконки</div>
                </div>
            </div>
        </div>
    )
}