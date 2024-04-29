import { For, useContext } from 'solid-js';
import { CursorContext } from '@/layouts/context/cursor';
import RouterAnimation from '@/components/RouterAnimation';
import { A } from '@solidjs/router';
import NavLine from '@/components/NavLine';

export default function Album() {
    const images = [
        'https://images.pexels.com/photos/6954544/pexels-photo-6954544.jpeg?auto=compress&cs=tinysrgb&w=1600',
        'https://images.pexels.com/photos/6954509/pexels-photo-6954509.jpeg?auto=compress&cs=tinysrgb&w=1600',
        'https://images.pexels.com/photos/11153461/pexels-photo-11153461.jpeg?auto=compress&cs=tinysrgb&w=1600',
        'https://images.pexels.com/photos/5049271/pexels-photo-5049271.jpeg?auto=compress&cs=tinysrgb&w=1600',
        'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
        'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
        'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
        'https://images.pexels.com/photos/19057526/pexels-photo-19057526.jpeg?auto=compress&cs=tinysrgb&w=1600',
    ];
    const { setActive, setInactive } = useContext(CursorContext);

    return (
        <div class="w-full h-full flex bg-black/90 z-top">
            <NavLine title="焦糖拿铁" subTitle={'Caramel latte'} class="left-36" position="center" />
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
                                        <small class="font-bold text-white/50 uppercase">Milk Company</small>
                                        <h5 class="text-2xl text-white font-bold uppercase">牛奶公司</h5>
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
