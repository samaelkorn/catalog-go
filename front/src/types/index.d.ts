interface TColor {
    id: number
    name: string
    code: string
}

interface TStatus {
    id: number
    name: string
    code: string
}

interface TProduct {
    id: number
    name: string
    image: string
    price: number
    color: TColor
    status: TStatus
}