/*
 * Copyright 2022 Cisco Systems, Inc. and its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import { useQuery } from 'react-query';
import Api from './api';

export namespace AnalyserData {
  export type Analyser = {
    status: string;
    name_id: string;
    title: string;
  };
}

export function useFetchAnalyzerList() {
  return useQuery('analyzer-list', () => Api.get('/analyzers?status=active'));
}
export async function TriggerReanalyze(
  serviceId: string,
  specId: string,
  analyzers: string[],
) {
  const url = `/services/${serviceId}/specs/${specId}/analyses`;
  const payload = {
    analyzers,
  };
  const result = await Api.post(url, payload);
  return result;
}