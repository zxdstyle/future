import type { CrudFilters } from "@refinedev/core"
import { mapOperator } from "./mapOperator"

export function generateFilter(filters?: CrudFilters) {
    const queryFilters: { [key: string]: string } = {}

    if (filters) {
        filters.map((filter) => {
            if (filter.operator === "or" || filter.operator === "and") {
                throw new Error(
                    // eslint-disable-next-line max-len
                    `[@refinedev/simple-rest]: \`operator: ${filter.operator}\` is not supported. You can create custom data provider. https://refine.dev/docs/api-reference/core/providers/data-provider/#creating-a-data-provider`,
                )
            }

            if ("field" in filter) {
                const { field, operator, value } = filter

                if (field === "q") {
                    queryFilters[field] = value
                    return
                }

                const [mappedOperator, val] = mapOperator(operator, value)
                queryFilters[`query[${field}]${mappedOperator}`] = val
            }
        })
    }

    return queryFilters
}
