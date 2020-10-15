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
 * @interface ControllersSystemequipment
 */
export interface ControllersSystemequipment {
    /**
     * 
     * @type {number}
     * @memberof ControllersSystemequipment
     */
    nameEquipmentID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersSystemequipment
     */
    physicianID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersSystemequipment
     */
    stockEquipmentID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersSystemequipment
     */
    typeEquipmentID?: number;
}

export function ControllersSystemequipmentFromJSON(json: any): ControllersSystemequipment {
    return ControllersSystemequipmentFromJSONTyped(json, false);
}

export function ControllersSystemequipmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersSystemequipment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'nameEquipmentID': !exists(json, 'nameEquipmentID') ? undefined : json['nameEquipmentID'],
        'physicianID': !exists(json, 'physicianID') ? undefined : json['physicianID'],
        'stockEquipmentID': !exists(json, 'stockEquipmentID') ? undefined : json['stockEquipmentID'],
        'typeEquipmentID': !exists(json, 'typeEquipmentID') ? undefined : json['typeEquipmentID'],
    };
}

export function ControllersSystemequipmentToJSON(value?: ControllersSystemequipment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'nameEquipmentID': value.nameEquipmentID,
        'physicianID': value.physicianID,
        'stockEquipmentID': value.stockEquipmentID,
        'typeEquipmentID': value.typeEquipmentID,
    };
}


