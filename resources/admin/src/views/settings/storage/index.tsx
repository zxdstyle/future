import { Badge, Button, Flex, Form, Image, Input, Modal, Tag, Typography, theme } from "antd"
import { useTable } from "@refinedev/antd"
import type { CrudFilter } from "@refinedev/core"
import { useDelete } from "@refinedev/core"
import { CreateStorage } from "@/components/FileManager"
import HDD from "@/assets/icons/hdd.png"

export function Component() {
    const { searchFormProps, tableProps } = useTable<Storage>({
        resource: "storages",
        pagination: { pageSize: 100 },
        sorters: {
            initial: [{ field: "id", order: "desc" }],
        },
        syncWithLocation: false,
        onSearch: (data: Record<string, any>) => {
            const filters: CrudFilter[] = []
            for (const dataKey in data) {
                filters.push({
                    field: dataKey,
                    operator: "contains",
                    value: data[dataKey],
                })
            }
            return filters
        },
    })

    const { token } = theme.useToken()
    const [modal, modalContext] = Modal.useModal()
    const { mutate, isLoading } = useDelete()
    const handleDelete = (id: number) => {
        modal.confirm({
            title: "删除位置",
            content: "删除存储位置时，只会从数据库中移除所有与之相关的文件数据，但是不会删除文件本身。",
            okType: "danger",
            onOk: async () => {
                mutate({ resource: "storages", id })
            },
        })
    }

    return (
        <div className="px-12 py-6 w-full">
            {modalContext}
            <Flex justify="space-between" align="center">
                <div>
                    <h1 className="text-white text-2xl mb-2">存储</h1>
                    <h2>管理您的存储位置</h2>
                </div>
                <Flex gap={12}>
                    <Form {...searchFormProps}>
                        <Form.Item noStyle name={["name"]}>
                            <Input prefix={<IconTablerSearch />} />
                        </Form.Item>
                    </Form>

                    <CreateStorage>
                        <Button type="primary">添加存储</Button>
                    </CreateStorage>
                </Flex>
            </Flex>

            <ul className="flex flex-col gap-4 mt-2">
                {tableProps?.dataSource?.map(item => (
                    <li
                        key={item.id}
                        className="px-4 py-2 rounded-lg border border-solid"
                        style={{ backgroundColor: token.colorBgContainer, borderColor: token.colorBorder }}
                    >
                        <Flex align="center" justify="space-between">
                            <Flex align="center" gap={12}>
                                <Image src={HDD} width={52} preview={false} />
                                <Flex vertical>
                                    <Typography.Title level={4} className="!text-white !m-0">{item.name}</Typography.Title>
                                    <Typography.Text className="text-sm">{item.option.local.folder}</Typography.Text>
                                </Flex>
                            </Flex>

                            <Flex align="center">
                                <Tag className="h-8 flex items-center" color={token.colorBgSpotlight} bordered={false}>
                                    <Badge color="green" className="mr-2" />
                                    已连接
                                </Tag>
                                <Button
                                    onClick={() => handleDelete(item.id)}
                                    className="border-none"
                                    style={{ backgroundColor: token.colorBgSpotlight }}
                                    icon={<IconPhTrash />}
                                    loading={isLoading}
                                />
                            </Flex>
                        </Flex>
                    </li>
                ))}
            </ul>

        </div>
    )
}
