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
    EntMedicalEquipment,
    EntMedicalEquipmentFromJSON,
    EntMedicalEquipmentFromJSONTyped,
    EntMedicalEquipmentToJSON,
    EntMedicalType,
    EntMedicalTypeFromJSON,
    EntMedicalTypeFromJSONTyped,
    EntMedicalTypeToJSON,
    EntPhysician,
    EntPhysicianFromJSON,
    EntPhysicianFromJSONTyped,
    EntPhysicianToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSystemequipmentEdges
 */
export interface EntSystemequipmentEdges {
    /**
     * 
     * @type {EntMedicalEquipment}
     * @memberof EntSystemequipmentEdges
     */
    medicalequipment?: EntMedicalEquipment;
    /**
     * 
     * @type {EntMedicalType}
     * @memberof EntSystemequipmentEdges
     */
    medicaltype?: EntMedicalType;
    /**
     * 
     * @type {EntPhysician}
     * @memberof EntSystemequipmentEdges
     */
    physician?: EntPhysician;
}

export function EntSystemequipmentEdgesFromJSON(json: any): EntSystemequipmentEdges {
    return EntSystemequipmentEdgesFromJSONTyped(json, false);
}

export function EntSystemequipmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSystemequipmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'medicalequipment': !exists(json, 'Medicalequipment') ? undefined : EntMedicalEquipmentFromJSON(json['Medicalequipment']),
        'medicaltype': !exists(json, 'Medicaltype') ? undefined : EntMedicalTypeFromJSON(json['Medicaltype']),
        'physician': !exists(json, 'Physician') ? undefined : EntPhysicianFromJSON(json['Physician']),
    };
}

export function EntSystemequipmentEdgesToJSON(value?: EntSystemequipmentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'medicalequipment': EntMedicalEquipmentToJSON(value.medicalequipment),
        'medicaltype': EntMedicalTypeToJSON(value.medicaltype),
        'physician': EntPhysicianToJSON(value.physician),
    };
}


