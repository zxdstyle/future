import type { ResourceProps } from "@refinedev/core"

const resources: ResourceProps[] = [
    {
        name: "mission-options",
        list: "/mission-options",
        create: "/mission-options/create",
        edit: "/mission-options/:id/edit",
        show: "/mission-options/:id",
        meta: {
            canDelete: false,
            label: "任务配置",
        },
    },
    {
        name: "dictionaries",
        list: "/dictionaries",
        // create: '/mission-options/create',
        // edit: '/mission-options/:id/edit',
        // show: '/mission-options/:id',
        meta: {
            canDelete: false,
            label: "数据字典",
        },
    },
    {
        name: "categories",
        list: "/categories",
        create: "/categories/create",
        edit: "/categories/edit/:id",
        show: "/categories/show/:id",
        meta: {
            canDelete: true,
        },
    },
]

export default resources
