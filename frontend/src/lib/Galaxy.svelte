<script>
    import { createEventDispatcher } from 'svelte';
    const dispatch = createEventDispatcher();

    import Planet from './Elements/Planet.svelte';
    import Mars from './Elements/Icons/Mars.svelte';
    import Mercury from './Elements/Icons/Mercury.svelte';
    import Venus from './Elements/Icons/Venus.svelte';
    import Earth from './Elements/Icons/Earth.svelte';
    import Jupiter from './Elements/Icons/Jupiter.svelte';
    import Saturn from './Elements/Icons/Saturn.svelte';
    import Uranus from './Elements/Icons/Uranus.svelte';
    import Neptune from './Elements/Icons/Neptune.svelte';

    // planets list
    let planets = [
        'mercury',
        'venus',
        'earth',
        'mars',
        'jupiter',
        'saturn',
        'uranus',
        'neptune',
    ];
</script>

<div class="relative bg-white py-8">
    <div class="lg:mx-auto lg:max-w-7xl lg:px-8 lg:items-center">
        <div class="relative sm:py-16 lg:py-0">
            <div
                class="relative mx-auto max-w-md px-4 sm:max-w-3xl sm:px-6 lg:px-0 lg:max-w-none lg:py-20"
            >
                <div
                    class="relative pb-10 h-full w-full rounded-2xl shadow-xl overflow-hidden"
                >
                    <div
                        class="absolute inset-0 bg-indigo-500 mix-blend-multiply"
                    ></div>
                    <div
                        class="absolute inset-0 bg-gradient-to-t from-indigo-600 via-indigo-600 opacity-90"
                    ></div>
                    <!--  -->
                    <div
                        class="relative mt-24 mb-24 flex justify-center items-center space-x-4"
                    >
                        <!-- -->
                        <Planet
                            name="{planets[0]}"
                            on:click="{() => {
                                dispatch('select', planets[0]);
                            }}"
                        >
                            <Mercury />
                        </Planet>
                        <Planet
                            name="{planets[1]}"
                            on:click="{() => {
                                dispatch('select', planets[1]);
                            }}"
                        >
                            <Venus />
                        </Planet>
                        <Planet
                            name="{planets[2]}"
                            on:click="{() => {
                                dispatch('select', planets[2]);
                            }}"
                        >
                            <Earth />
                        </Planet>
                        <Planet
                            name="{planets[3]}"
                            on:click="{() => {
                                dispatch('select', planets[3]);
                            }}"
                        >
                            <Mars />
                        </Planet>
                        <Planet
                            name="{planets[4]}"
                            on:click="{() => {
                                dispatch('select', planets[4]);
                            }}"
                        >
                            <Jupiter />
                        </Planet>
                        <Planet
                            name="{planets[5]}"
                            on:click="{() => {
                                dispatch('select', planets[5]);
                            }}"
                        >
                            <Saturn />
                        </Planet>
                        <Planet
                            name="{planets[6]}"
                            on:click="{() => {
                                dispatch('select', planets[6]);
                            }}"
                        >
                            <Uranus />
                        </Planet>
                        <Planet
                            name="{planets[7]}"
                            on:click="{() => {
                                dispatch('select', planets[7]);
                            }}"
                        >
                            <Neptune />
                        </Planet>
                        <!--  -->
                    </div>
                    <!--  -->
                    <div
                        class="relative flex justify-center items-center px-8 mt-8 space-x-5"
                    >
                        <blockquote>
                            <div
                                class="relative text-lg font-medium text-white md:flex-grow"
                            >
                                <svg
                                    class="absolute top-0 left-0 transform -translate-x-3 -translate-y-2 h-8 w-8 text-indigo-500"
                                    fill="currentColor"
                                    viewBox="0 0 32 32"
                                    aria-hidden="true"
                                >
                                    <path
                                        d="M9.352 4C4.456 7.456 1 13.12 1 19.36c0 5.088 3.072 8.064 6.624 8.064 3.36 0 5.856-2.688 5.856-5.856 0-3.168-2.208-5.472-5.088-5.472-.576 0-1.344.096-1.536.192.48-3.264 3.552-7.104 6.624-9.024L9.352 4zm16.512 0c-4.8 3.456-8.256 9.12-8.256 15.36 0 5.088 3.072 8.064 6.624 8.064 3.264 0 5.856-2.688 5.856-5.856 0-3.168-2.304-5.472-5.184-5.472-.576 0-1.248.096-1.44.192.48-3.264 3.456-7.104 6.528-9.024L25.864 4z"
                                    ></path>
                                </svg>
                                <p
                                    class="relative font-semibold text-3xl text-gray-100"
                                >
                                    How fast is the milky way moving?
                                </p>
                            </div>
                        </blockquote>
                        <!--  -->
                        <footer>
                            <div
                                class="flex justify-center items-center py-3 w-full"
                            >
                                <button
                                    on:click="{() => {
                                        dispatch('check');
                                    }}"
                                    class="py-2 px-5 rounded text-lg font-semibold text-green-800 bg-green-400 hover:bg-green-600"
                                >
                                    find out
                                </button>
                            </div>
                        </footer>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
