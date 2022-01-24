export default class EventService {
    
    // credits: https://www.webtips.dev/intro-to-design-patterns-create-your-own-pubsub-library
    
    subscriptions = {};

    /**
     * @method publish
     *
     * @param event { string } - name of the event to publish
     * @param data  { any }    - optional data to pass on to subscription
     *
     * @example
     * - event.publish('event');
     * - event.publish('event', { data: 'customData' });
     * - event.publish('event event1 event2', { data: 'customData' });
     */
    publish(event, data) {
        for (const key in this.subscriptions) {
            if (this.subscriptions.hasOwnProperty(key)) {
                if (key === event) {
                    for (const currentInstance of this.subscriptions[key]) {
                        currentInstance.callback(data);

                        const evt = new CustomEvent(event, {'detail': data } );
                        window.dispatchEvent(evt);
                    }
                }
            }
        }
    }

    /**
     * @method subscribe
     *
     * @param event    { string } 	- name of the event to subscribe to
     * @param callback { function } - function callback to execute once the event has been published
     *
     * @example
     * - event.subscribe('event', function () { ... });
     * - event.subscribe('event event1 event2', function (data) { ... });
     * - event.subscribe('event.namespaced', function (data) { ... });
     */
    subscribe(event, callback) {
        const events = event.split(' ');
        let eventArray = [];

        if (!callback) {
            throw new Error('callback is missing from subscribe: ' + event);
        }

        for (const singleEvent of events) {
            if (singleEvent.indexOf('.') > -1) {
                eventArray = singleEvent.split('.');
            }

            if (!this.subscriptions[eventArray[0] || singleEvent]) {
                this.subscriptions[eventArray[0] || singleEvent] = [];
            }

            for (const currentInstance of this.subscriptions[eventArray[0] || singleEvent]) {
                if (currentInstance.namespace === eventArray[1]) {
                    return;
                }
            }

            this.subscriptions[eventArray[0] || singleEvent].push({
                callback,
                namespace: eventArray[1]
            });

            eventArray = [];
        }
    }

    /**
     * @method unsubscribe
     *
     * @param event { string } - name of the event to unsubscribe from
     *
     * @example
     * - event.unsubscribe('event');
     * - event.unsubscribe('event event1 event2');
     * - event.unsubscribe('event.namespaced');
     */
    unsubscribe(event) {
        const events = event.split(' ');
        let eventArray;

        for (const currentEvent of events) {
            eventArray = currentEvent.split('.') || currentEvent;

            if (this.subscriptions[eventArray[0]]) {
                for (let j = 0; j < this.subscriptions[eventArray[0]].length; j++) {
                    if (eventArray[1]) {

                        if (this.subscriptions[eventArray[0]][j].namespace === eventArray[1]) {
                            this.subscriptions[eventArray[0]].splice(j, 1);
                        }

                        if (!this.subscriptions[eventArray[0]].length) {
                            delete this.subscriptions[eventArray[0]];
                            break;
                        }

                    } else {
                        delete this.subscriptions[eventArray[0]];
                        break;
                    }
                }
            }
        }

    };

}