import { createMasonry } from '@solid-primitives/masonry';
import { createSignal, useContext } from 'solid-js';
import { createBreakpoints } from '@solid-primitives/media';
import { createElementSize } from '@solid-primitives/resize-observer';
import { CursorContext } from '@/layouts/context/cursor';
import SmoothScroll from '@/components/SmoothScroll';
import RouterAnimation from '@/components/RouterAnimation';

interface Props {
    images?: string[];
}

const getRandomHeight = () => Math.floor(Math.random() * 300) + 100;
export default function Masonry(props: Props) {
    const [items, setItems] = createSignal([
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/22609359/pexels-photo-22609359.jpeg?auto=compress&cs=tinysrgb&w=800&lazy=load',
            width: 466,
            height: 312,
        },
        {
            src: 'https://images.pexels.com/photos/22601821/pexels-photo-22601821.jpeg?auto=compress&cs=tinysrgb&w=800&lazy=load',
            width: 466,
            height: 312,
        },
        {
            src: 'https://images.pexels.com/photos/22662125/pexels-photo-22662125.jpeg?auto=compress&cs=tinysrgb&w=800&lazy=load',
            width: 466,
            height: 583,
        },
        {
            src: 'https://images.pexels.com/photos/22643799/pexels-photo-22643799.jpeg?auto=compress&cs=tinysrgb&w=800&lazy=load',
            width: 466,
            height: 312,
        },
        {
            src: 'https://images.pexels.com/photos/22643799/pexels-photo-22643799.jpeg?auto=compress&cs=tinysrgb&w=800&lazy=load',
            width: 466,
            height: 312,
        },
        {
            src: 'https://images.pexels.com/photos/22643799/pexels-photo-22643799.jpeg?auto=compress&cs=tinysrgb&w=800&lazy=load',
            width: 466,
            height: 312,
        },
        {
            src: 'https://images.pexels.com/photos/22643799/pexels-photo-22643799.jpeg?auto=compress&cs=tinysrgb&w=800&lazy=load',
            width: 466,
            height: 312,
        },
    ]);

    const br = createBreakpoints({
        sm: '640px',
        md: '768px',
        lg: '1024px',
        xl: '1280px',
    });
    const { setActive, setInactive } = useContext(CursorContext);

    const masonry = createMasonry({
        source: items,
        columns: () => {
            return 3;
        },
        mapHeight(item) {
            // observe the height of the element
            // return the accessor of the height of the element
            return () => item.height + 100;
        },
        mapElement(item, i) {
            return (
                <div
                    class="cursor-pointer overflow-hidden group rounded"
                    style={{
                        width: `${item.source.width}px`,
                        height: `${item.margin() + item.source.height}px`,
                        order: item.order(),
                    }}
                    onMouseOver={setActive}
                    onMouseLeave={setInactive}
                >
                    <img class="group-hover:scale-105 transition-all duration-500" src={item.source.src} alt="" />
                </div>
            );
        },
    });

    return (
        <SmoothScroll>
            <section class="px-44 py-32 w-full h-full">
                <RouterAnimation
                    initial={{ opacity: 0, y: 100 }}
                    animate={{ opacity: 1, y: 0 }}
                    exit={{ opacity: 0, y: -100, transition: { duration: 1 } }}
                    transition={{ duration: 2 }}
                    class="mx-32 h-full"
                >
                    <p class="mx-auto max-w-4xl text-center mb-2">
                        在牛奶公司中，您不仅可以品尝到榛果巧克力风味，还可以在奶咖咖中喝到热带水果的香气，尾端还有浓郁的可可风味。
                        风味：榛果、热带水果、可可、凤梨、热带水果香气、榛果巧克力、浓郁可可
                    </p>
                    <div class="box-border flex min-h-screen w-full flex-col items-center justify-center space-y-4 p-24 text-white overflow-scroll mx-auto">
                        <div
                            class="flex flex-col flex-wrap justify-start gap-10 "
                            style={{ height: `${masonry.height() - 24}px` }}
                        >
                            {masonry()}
                        </div>
                    </div>
                </RouterAnimation>
            </section>
        </SmoothScroll>
    );
}
