// import {Module} from '../module/Module'
//
// class ModuleCheckout implements Module{
//     path: string;
//     component: () => Promise<{ default: any }>;
//     constructor(path: string, component: () => Promise<{ default: any }>) {
//         this.path = path;
//         this.component = component;
//     }
// }
//
// // Instanciando a classe e exportando
// // @ts-ignore
// export const checkoutModule = new ModuleCheckout("/checkout", () => import("./checkout.svelte"));