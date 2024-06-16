import { For, useContext } from 'solid-js';
import { CursorContext } from '@/layouts/context/cursor';
import RouterAnimation from '@/components/RouterAnimation';
import { A } from '@solidjs/router';
import NavLine from '@/components/NavLine';
import Cover from '@/assets/images/baby/IMG_0073-2.jpg.webp';

export default function Album() {
    const images = [Cover];
    const { setActive, setInactive } = useContext(CursorContext);

    return (
        <div class="w-full h-full flex bg-black/90 z-top">
            <NavLine title="百日照" subTitle={'Caramel latte'} class="left-36" position="center" />
            <section class="px-44 py-32 w-full h-full overflow-auto">
                <RouterAnimation
                    initial={{ opacity: 0, y: 100 }}
                    animate={{ opacity: 1, y: 0 }}
                    exit={{ opacity: 0, y: -100, transition: { duration: 1 } }}
                    transition={{ duration: 2 }}
                    class="mx-32"
                >
                    <ul class="grid grid-cols-3 gap-12">
                        <For each={images}>
                            {item => (
                                <li class="group" onMouseOver={() => setActive()} onMouseLeave={() => setInactive()}>
                                    <A href={'/album/detail'}>
                                        <div class="rounded overflow-hidden">
                                            <img
                                                src={item}
                                                class="group-hover:scale-105 transition-all duration-500"
                                                alt=""
                                            />
                                        </div>
                                        <small class="font-bold text-white/50 uppercase">Hundred Days</small>
                                        <h5 class="text-2xl text-white font-bold uppercase">百日照</h5>
                                    </A>
                                </li>
                            )}
                        </For>
                    </ul>
                </RouterAnimation>
            </section>
            <NavLine title="返回" subTitle="return" class="right-36" position="bottom" onClick={() => history.back()} />
        </div>
    );
}
