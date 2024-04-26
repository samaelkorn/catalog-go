const envApi = process.env.DATA_API_URL! ?? '/'

export const urlApi = (...arg: string[]): string => `${envApi}/${arg.join('/')}`