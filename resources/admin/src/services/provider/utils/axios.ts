import type { HttpError } from "@refinedev/core"
import axios from "axios"

const axiosInstance = axios.create()

axiosInstance.interceptors.response.use(
    (response) => {
        return response.data
    },
    (error) => {
        const customError: HttpError = {
            ...error,
            message: error.response?.data?.msg,
            statusCode: error.response?.status,
        }
        return Promise.reject(customError)
    },
)

export { axiosInstance }
