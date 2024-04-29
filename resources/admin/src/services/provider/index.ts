import type { DataProvider } from "@refinedev/core"
import type { AxiosInstance } from "axios"
import stringify from "query-string"
import { axiosInstance, generateFilter, generateSort } from "./utils"
import type { ListResponse } from "@/services/provider/types/response"

type MethodTypes = "get" | "delete" | "head" | "options"
type MethodTypesWithBody = "post" | "put" | "patch"

export function dataProvider(
    apiUrl: string,
    httpClient: AxiosInstance = axiosInstance,
): Omit<Required<DataProvider>, "createMany" | "updateMany" | "deleteMany"> {
    return {
        getList: async ({ resource, pagination, filters, sorters, meta }) => {
            const url = `${apiUrl}/${resource}`

            const { current = 1, pageSize = 10, mode = "server" } = pagination ?? {}

            const { headers: headersFromMeta, method, with: relations } = meta ?? {}
            const requestMethod = (method as MethodTypes) ?? "get"

            const queryFilters = generateFilter(filters)

            const query: {
                per_page?: number
                page?: number
                order_by?: string
                order?: string
                with?: string
            } = {}

            if (mode === "server") {
                query.per_page = pageSize
                query.page = current
            }

            const generatedSort = generateSort(sorters)
            if (generatedSort) {
                const { _sort, _order } = generatedSort
                query.order_by = _sort.join(",")
                query.order = _order.join(",")
            }
            if (relations)
                query.with = relations.join(",")

            const combinedQuery = { ...query, ...queryFilters }
            const urlWithQuery = Object.keys(combinedQuery).length ? `${url}?${stringify.stringify(combinedQuery)}` : url

            const { data, meta: metas } = await httpClient[requestMethod]<any, ListResponse<any>>(urlWithQuery, {
                headers: headersFromMeta,
            })

            return {
                data,
                total: metas?.pagination?.total,
            }
        },

        getMany: async ({ resource, ids, meta }) => {
            const { headers, method } = meta ?? {}
            const requestMethod = (method as MethodTypes) ?? "get"

            const { data } = await httpClient[requestMethod](`${apiUrl}/${resource}?${stringify.stringify({ id: ids })}`, {
                headers,
            })

            return {
                data,
            }
        },

        create: async ({ resource, variables, meta }) => {
            const url = `${apiUrl}/${resource}`

            const { headers, method } = meta ?? {}
            const requestMethod = (method as MethodTypesWithBody) ?? "post"

            const { data } = await httpClient[requestMethod](url, variables, {
                headers,
            })

            return {
                data,
            }
        },

        update: async ({ resource, id, variables, meta }) => {
            const url = `${apiUrl}/${resource}/${id}`

            const { headers, method } = meta ?? {}
            const requestMethod = (method as MethodTypesWithBody) ?? "patch"

            const { data } = await httpClient[requestMethod](url, variables, {
                headers,
            })

            return {
                data,
            }
        },

        getOne: async ({ resource, id, meta }) => {
            const url = `${apiUrl}/${resource}/${id}`

            const { headers, method } = meta ?? {}
            const requestMethod = (method as MethodTypes) ?? "get"

            const { data } = await httpClient[requestMethod](url, { headers })

            return {
                data,
            }
        },

        deleteOne: async ({ resource, id, variables, meta }) => {
            const url = `${apiUrl}/${resource}/${id}`

            const { headers, method } = meta ?? {}
            const requestMethod = (method as MethodTypesWithBody) ?? "delete"

            const { data } = await httpClient[requestMethod](url, { data: variables, headers })

            return {
                data,
            }
        },

        getApiUrl: () => {
            return apiUrl
        },

        custom: async ({ url, method, filters, sorters, payload, query, headers }) => {
            let requestUrl = `${url}?`

            if (sorters) {
                const generatedSort = generateSort(sorters)
                if (generatedSort) {
                    const { _sort, _order } = generatedSort
                    const sortQuery = {
                        order_by: _sort.join(","),
                        order: _order.join(","),
                    }
                    requestUrl = `${requestUrl}&${stringify.stringify(sortQuery)}`
                }
            }

            if (filters) {
                const filterQuery = generateFilter(filters)
                requestUrl = `${requestUrl}&${stringify.stringify(filterQuery)}`
            }

            if (query)
                requestUrl = `${requestUrl}&${stringify.stringify(query)}`

            let axiosResponse
            switch (method) {
            case "put":
            case "post":
            case "patch":
                axiosResponse = await httpClient[method](url, payload, {
                    headers,
                })
                break
            case "delete":
                axiosResponse = await httpClient.delete(url, {
                    data: payload,
                    headers,
                })
                break
            default:
                axiosResponse = await httpClient.get(requestUrl, {
                    headers,
                })
                break
            }

            const { data } = axiosResponse

            return Promise.resolve({ data })
        },
    }
}

export default dataProvider("/api/v1")
