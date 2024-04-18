export default function Filter() {
    return (
        <div className="shadow-lg rounded-[10px] py-10 px-5 bg-white">
            <div className="flex">
                <div className="mr-2">
                    <select defaultValue="" className="border border-stone-400 text-sm rounded-[7px] p-2.5 pr-4">
                        <option selected value="">Выберите цвет</option>
                        <option value="US">United States</option>
                        <option value="CA">Canada</option>
                        <option value="FR">France</option>
                        <option value="DE">Germany</option>
                    </select>
                </div>
                <div>
                    <select defaultValue="" className="border border-stone-400 text-sm rounded-[7px] p-2.5 pr-4">
                        <option selected value="">Выберите статус</option>
                        <option value="US">United States</option>
                        <option value="CA">Canada</option>
                        <option value="FR">France</option>
                        <option value="DE">Germany</option>
                    </select>
                </div>
            </div>
        </div>
    )
}