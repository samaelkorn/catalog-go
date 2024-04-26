export default function Sort() {
    return (
        <div className="mt-10 shadow-lg rounded-[10px] py-7 px-5 bg-slate-100">
            <div>
                <select name="sort" className="bg-slate-100" defaultValue="">
                    <option value="">Сортировка по</option>
                    <option value="price">По цене</option>
                    <option value="name">По названию</option>
                </select>
            </div>
        </div>
    )
}