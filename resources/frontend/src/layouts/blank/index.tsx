import Cursor from './component/Cursor';
import { JSX } from 'solid-js';
import Header from './component/Header';
import Footer from './component/Footer';
import { RouterStateProvider } from '@/layouts/context/router';

interface IndexProps {
    children?: JSX.Element;
}

export default function Index(props: IndexProps) {
    return (
        <RouterStateProvider>
            <Cursor>
                <Header />

                <main class="w-full h-full">{props.children}</main>

                <Footer />
            </Cursor>
        </RouterStateProvider>
    );
}
