import { Flex } from "antd"
import useDataSource from "../hooks/useDataSource"
import Footer from "./Footer"
import Overview from "./Overview"
import GridView from "./GridView"
import Navbar from "./Navbar"
import { Provider } from "./context"

interface ContentProps {
    root: string
}

export default function ({ root }: ContentProps) {
    return (
        <Provider root={root}>
            <Content />
        </Provider>
    )
}

function Content() {
    const { data } = useDataSource()

    return (
        <Flex vertical className="h-full w-full">
            <Navbar />

            <div className="h-full flex-1 overflow-hidden relative">
                <GridView data={data} />

                <Overview />
            </div>

            <Footer />
        </Flex>
    )
}
