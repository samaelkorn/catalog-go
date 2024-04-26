import Filter from './Filter'
import Sort from './Sort'
import Catalog from './Catalog'
import { getProducts, getColors, getStatuses } from '@/api'

export default async function CatalogPage() {
    const { data }: { data: TProduct[] } = await getProducts()
    const { data: colors }: { data: TColor[] } = await getColors()
    const { data: statuses }: { data: TStatus[] } = await getStatuses()

    return (
        <div>
            <div className="container mx-auto">
                <div>
                    <div className="my-10">
                        <h1 className="text-3xl font-medium">Новые мотоциклы</h1>
                    </div>
                </div>
            </div>
            <div className="bg-slate-100 py-8">
                <div className="container mx-auto">
                    <Filter colors={colors} statuses={statuses} />
                </div>
            </div>
            <div className="container mx-auto">
                <Sort />
            </div>
            <div className="container mx-auto">
                <Catalog data={data} />
            </div>
        </div>
    )
}