import Footer from "./Footer"
import Navbar from "./Navbar"
import useDataSource from "./hooks/useDataSource"
import GridView from "./GridView"
import Overview from "./Overview"

export default function Manager() {
    const { data } = useDataSource()

    return (
        <section className="flex flex-col h-full">
            <Navbar />

            <div className="h-full flex-1 overflow-hidden relative">
                <GridView data={data} />

                <Overview />
            </div>

            <Footer />
        </section>
    )
}
