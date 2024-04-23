import Filter from './Filter'
import Sort from './Sort'
import Catalog from './Catalog'

export default function CatalogPage() {
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
                    <Filter />
                </div>
            </div>
            <div className="container mx-auto">
                <Sort />
            </div>
            <div className="container mx-auto">
                <Catalog />
            </div>
        </div>
    )
}