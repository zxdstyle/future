import GridItem from "@/components/Storage/GridView/GridItem.tsx"

interface GridViewProps {
    data?: FileDescription[]
}

export default function GridView({ data = [] }: GridViewProps) {
    return (
        <ul className="flex flex-wrap gap-x-6 gap-y-4 h-full overflow-auto content-start">
            {data.map(item => <GridItem item={item} key={item.filename} />)}
        </ul>
    )
}
