const envNext = process.env.DATA_NEXT_URL! ?? 'http://localhost:3000'

export const urlNext = (...arg: string[]): string => `${envNext}/${arg.join('/')}`