const envNext = process.env.DATA_NEXT_URL! ?? '/'

export const urlNext = (...arg: string[]): string => `${envNext}/${arg.join('/')}`