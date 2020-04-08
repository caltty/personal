export function getServerConfig() {
    return Promise.resolve({
        data: {
            port: 4377,
            certFile: "/var/1.cert",
            sslKey: "123456",
        }
    })
}