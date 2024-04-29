import type { CrudOperators } from "@refinedev/core"

export function mapOperator(operator: CrudOperators, value: any): [string, any] {
    switch (operator) {
    case "ne":
    case "gte":
    case "lte":
        return [`_${operator}`, value]
    case "contains":
        return ["_like", value]
    case "eq":
    default:
        return ["", value]
    }
}
