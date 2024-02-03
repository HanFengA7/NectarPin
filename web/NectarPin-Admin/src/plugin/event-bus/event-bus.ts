// event-bus.ts
import mitt from 'mitt';

// 创建一个全局事件总线
const eventBus = mitt();

export default eventBus;
