/**
 * Penlook Project
 *
 * Copyright (c) 2015 Penlook Development Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Authors:
 *     Loi Nguyen       <loint@penlook.com>
 *     Nam Vo           <namvh@penlook.com>
 */

library test.transport;

import 'package:unittest/unittest.dart';
import 'package:unittest/html_config.dart';
import 'package:socket/polling.dart';
import 'package:socket/option.dart';

// Test abstract class
class TestTransport extends Transport {}

class Context {
    String test;
}

void main() {

    useHtmlConfiguration();

    test("transport synchronous", () {

        var test = new TestTransport();
        var context = new Context();
        var option  = new Option();

        test.syncRequest(context, option, (Context context, Map<String, Map> response) {

        });

    });

    test("transport asynchronous", () {

    });

    test("transport request", () {

    });

}