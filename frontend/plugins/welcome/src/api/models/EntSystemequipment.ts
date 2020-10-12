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
import {
    EntSystemequipmentEdges,
    EntSystemequipmentEdgesFromJSON,
    EntSystemequipmentEdgesFromJSONTyped,
    EntSystemequipmentEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSystemequipment
 */
export interface EntSystemequipment {
    /**
     * AddedTime holds the value of the "added_time" field.
     * @type {string}
     * @memberof EntSystemequipment
     */
    addedTime?: string;
    /**
     * 
     * @type {EntSystemequipmentEdges}
     * @memberof EntSystemequipment
     */
    edges?: EntSystemequipmentEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSystemequipment
     */
    id?: number;
}

export function EntSystemequipmentFromJSON(json: any): EntSystemequipment {
    return EntSystemequipmentFromJSONTyped(json, false);
}

export function EntSystemequipmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSystemequipment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addedTime': !exists(json, 'added_time') ? undefined : json['added_time'],
        'edges': !exists(json, 'edges') ? undefined : EntSystemequipmentEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntSystemequipmentToJSON(value?: EntSystemequipment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added_time': value.addedTime,
        'edges': EntSystemequipmentEdgesToJSON(value.edges),
        'id': value.id,
    };
}

