/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface EntHooks
 */
export interface EntHooks {
    /**
     * 
     * @type {Array<object>}
     * @memberof EntHooks
     */
    user?: Array<object>;
}

export function EntHooksFromJSON(json: any): EntHooks {
    return EntHooksFromJSONTyped(json, false);
}

export function EntHooksFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntHooks {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'user': !exists(json, 'user') ? undefined : json['user'],
    };
}

export function EntHooksToJSON(value?: EntHooks | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'user': value.user,
    };
}


