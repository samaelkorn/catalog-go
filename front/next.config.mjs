/** @type {import('next').NextConfig} */
const nextConfig = {
    sassOptions: {
        includePaths: ['./src/styles'],
    },
    images: {
        remotePatterns: [
            {
                protocol: process.env.DATA_IMG_PROTOCOL,
                hostname: process.env.DATA_IMG_HOSTNAME,
                port: process.env.DATA_IMG_PORT,
                pathname: '/storage/img/**',
            },
        ],
    },
}

export default nextConfig
