import { createSignal, For, onMount, useContext } from 'solid-js';
import { CursorContext } from '@/layouts/context/cursor';
import RouterAnimation from '@/components/RouterAnimation';
import { scroll } from 'motion';
import { Motion } from 'solid-motionone';
import ProgressBar from '@/components/ProgressBar';

export default function Carousel() {
    const [items, setItems] = createSignal([
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
        {
            src: 'https://images.pexels.com/photos/20216289/pexels-photo-20216289.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2',
            width: 467,
            height: 621,
        },
    ]);

    const { setActive, setInactive } = useContext(CursorContext);

    let container: HTMLDivElement | undefined = undefined;

    const [percent, setPercent] = createSignal(0);
    onMount(() => {
        scroll(({ x }) => setPercent(x.progress), { container, axis: 'x' });
    });

    return (
        <section class="py-32 w-full h-full">
            <RouterAnimation
                initial={{ opacity: 0, y: 100 }}
                animate={{ opacity: 1, y: 0 }}
                exit={{ opacity: 0, y: -100, transition: { duration: 1 } }}
                transition={{ duration: 2 }}
                class="h-full w-full"
            >
                <div class="flex w-full h-full flex-col items-center justify-between mx-auto gap-12">
                    <p class="mx-auto max-w-4xl text-center mb-2 text-gray">
                        在牛奶公司中，您不仅可以品尝到榛果巧克力风味，还可以在奶咖咖中喝到热带水果的香气，尾端还有浓郁的可可风味。
                        风味：榛果、热带水果、可可、凤梨、热带水果香气、榛果巧克力、浓郁可可
                    </p>

                    <div class="w-full h-full flex overflow-x-scroll gap-x-10  p-0" ref={container}>
                        <For each={items()}>
                            {item => (
                                <div
                                    class="flex-shrink-0  first:pl-60 last:pr-60 group  h-full overflow-hidden"
                                    onMouseOver={setActive}
                                    onMouseLeave={setInactive}
                                >
                                    <div
                                        style={{ 'background-image': `url(${item.src})` }}
                                        class="group-hover:scale-105 h-full w-96 transition-all duration-500 bg-no-repeat bg-contain bg-center"
                                    />
                                </div>
                            )}
                        </For>
                    </div>

                    <ProgressBar percent={percent()} class="w-3/5" />
                </div>
            </RouterAnimation>
        </section>
    );
}
