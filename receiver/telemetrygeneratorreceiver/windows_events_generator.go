// Copyright observIQ, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package telemetrygeneratorreceiver //import "github.com/observiq/bindplane-otel-collector/receiver/telemetrygeneratorreceiver"

import (
	"go.uber.org/zap"
)

// windowsEventsMetricsGenerator is a generator for Windows Event Log metrics. It generates a sampling of Windows Event Log metrics
// emulating the Windows Event Log receiver: github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver
func newWindowsEventsGenerator(logger *zap.Logger) (logGenerator, error) {
	return newOtlpGenerator(defaultWindowsEventsConfig, logger)
}

var defaultWindowsEventsConfig = GeneratorConfig{
	Type: "otlp",
	AdditionalConfig: map[string]any{
		"telemetry_type": "logs",
		"otlp_json":      `{"resourceLogs":[{"resource":{"attributes":[{"key":"host.name","value":{"stringValue":"winhost-windows"}},{"key":"os.type","value":{"stringValue":"windows"}}]},"scopeLogs":[{"scope":{},"logRecords":[{"timeUnixNano":"1710783558744839700","observedTimeUnixNano":"1710775543763436600","severityNumber":9,"severityText":"INFO","body":{"kvlistValue":{"values":[{"key":"execution","value":{"kvlistValue":{"values":[{"key":"process_id","value":{"intValue":"0"}},{"key":"thread_id","value":{"intValue":"0"}}]}}},{"key":"computer","value":{"stringValue":"winhost-windows"}},{"key":"message","value":{"stringValue":"Offline downlevel migration succeeded."}},{"key":"channel","value":{"stringValue":"Application"}},{"key":"task","value":{"stringValue":"0"}},{"key":"opcode","value":{"stringValue":"0"}},{"key":"event_data","value":{"kvlistValue":{}}},{"key":"event_id","value":{"kvlistValue":{"values":[{"key":"qualifiers","value":{"intValue":"49152"}},{"key":"id","value":{"intValue":"16394"}}]}}},{"key":"record_id","value":{"intValue":"24437"}},{"key":"level","value":{"stringValue":"Information"}},{"key":"keywords","value":{"arrayValue":{"values":[{"stringValue":"Classic"}]}}},{"key":"provider","value":{"kvlistValue":{"values":[{"key":"name","value":{"stringValue":"Microsoft-Windows-Security-SPP"}},{"key":"guid","value":{"stringValue":"{E23B33B0-C8C9-472C-A5F9-F2BDFEA0F156}"}},{"key":"event_source","value":{"stringValue":"Software Protection Platform Service"}}]}}},{"key":"system_time","value":{"stringValue":"2024-03-18T15:25:43.4767628Z"}}]}},"attributes":[{"key":"log_type","value":{"stringValue":"windows_event.application"}}],"traceId":"","spanId":""}]}]},{"resource":{"attributes":[{"key":"host.name","value":{"stringValue":"winhost-windows"}},{"key":"os.type","value":{"stringValue":"windows"}}]},"scopeLogs":[{"scope":{},"logRecords":[{"timeUnixNano":"1710783558634596100","observedTimeUnixNano":"1710775544762121700","severityNumber":9,"severityText":"INFO","body":{"kvlistValue":{"values":[{"key":"provider","value":{"kvlistValue":{"values":[{"key":"name","value":{"stringValue":"Service Control Manager"}},{"key":"guid","value":{"stringValue":"{555908d1-a6d7-4695-8e1e-26931d2012f4}"}},{"key":"event_source","value":{"stringValue":"Service Control Manager"}}]}}},{"key":"task","value":{"stringValue":"0"}},{"key":"execution","value":{"kvlistValue":{"values":[{"key":"thread_id","value":{"intValue":"3308"}},{"key":"process_id","value":{"intValue":"640"}}]}}},{"key":"record_id","value":{"intValue":"3141"}},{"key":"keywords","value":{"arrayValue":{"values":[{"stringValue":"Classic"}]}}},{"key":"event_data","value":{"kvlistValue":{"values":[{"key":"binary","value":{"stringValue":"7300700070007300760063002F0034000000"}},{"key":"data","value":{"arrayValue":{"values":[{"kvlistValue":{"values":[{"key":"param1","value":{"stringValue":"Software Protection"}}]}},{"kvlistValue":{"values":[{"key":"param2","value":{"stringValue":"running"}}]}}]}}}]}}},{"key":"event_id","value":{"kvlistValue":{"values":[{"key":"qualifiers","value":{"intValue":"16384"}},{"key":"id","value":{"intValue":"7036"}}]}}},{"key":"system_time","value":{"stringValue":"2024-03-18T15:25:43.3665192Z"}},{"key":"channel","value":{"stringValue":"System"}},{"key":"level","value":{"stringValue":"Information"}},{"key":"message","value":{"stringValue":"The Software Protection service entered the running state."}},{"key":"opcode","value":{"stringValue":"0"}},{"key":"computer","value":{"stringValue":"winhost-windows"}}]}},"attributes":[{"key":"log_type","value":{"stringValue":"windows_event.system"}}],"traceId":"","spanId":""}]}]},{"resource":{"attributes":[{"key":"host.name","value":{"stringValue":"winhost-windows"}},{"key":"os.type","value":{"stringValue":"windows"}}]},"scopeLogs":[{"scope":{},"logRecords":[{"timeUnixNano":"1710783574513483900","observedTimeUnixNano":"1710775560745899400","severityNumber":9,"severityText":"INFO","body":{"kvlistValue":{"values":[{"key":"system_time","value":{"stringValue":"2024-03-18T15:25:59.2454070Z"}},{"key":"level","value":{"stringValue":"Information"}},{"key":"message","value":{"stringValue":"An account failed to log on."}},{"key":"event_data","value":{"kvlistValue":{"values":[{"key":"data","value":{"arrayValue":{"values":[{"kvlistValue":{"values":[{"key":"SubjectUserSid","value":{"stringValue":"S-1-0-0"}}]}},{"kvlistValue":{"values":[{"key":"SubjectUserName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"SubjectDomainName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"SubjectLogonId","value":{"stringValue":"0x0"}}]}},{"kvlistValue":{"values":[{"key":"TargetUserSid","value":{"stringValue":"S-1-0-0"}}]}},{"kvlistValue":{"values":[{"key":"TargetUserName","value":{"stringValue":"Administrator"}}]}},{"kvlistValue":{"values":[{"key":"TargetDomainName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"Status","value":{"stringValue":"0xc000006d"}}]}},{"kvlistValue":{"values":[{"key":"FailureReason","value":{"stringValue":"%%2313"}}]}},{"kvlistValue":{"values":[{"key":"SubStatus","value":{"stringValue":"0xc000006a"}}]}},{"kvlistValue":{"values":[{"key":"LogonType","value":{"stringValue":"3"}}]}},{"kvlistValue":{"values":[{"key":"LogonProcessName","value":{"stringValue":"NtLmSsp "}}]}},{"kvlistValue":{"values":[{"key":"AuthenticationPackageName","value":{"stringValue":"NTLM"}}]}},{"kvlistValue":{"values":[{"key":"WorkstationName","value":{"stringValue":"workstation"}}]}},{"kvlistValue":{"values":[{"key":"TransmittedServices","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"LmPackageName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"KeyLength","value":{"stringValue":"0"}}]}},{"kvlistValue":{"values":[{"key":"ProcessId","value":{"stringValue":"0x0"}}]}},{"kvlistValue":{"values":[{"key":"ProcessName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"IpAddress","value":{"stringValue":"38.68.52.197"}}]}},{"kvlistValue":{"values":[{"key":"IpPort","value":{"stringValue":"0"}}]}}]}}}]}}},{"key":"execution","value":{"kvlistValue":{"values":[{"key":"process_id","value":{"intValue":"660"}},{"key":"thread_id","value":{"intValue":"3244"}}]}}},{"key":"computer","value":{"stringValue":"winhost-windows"}},{"key":"record_id","value":{"intValue":"107866"}},{"key":"task","value":{"stringValue":"Logon"}},{"key":"keywords","value":{"arrayValue":{"values":[{"stringValue":"Audit Failure"}]}}},{"key":"details","value":{"kvlistValue":{"values":[{"key":"Network Information","value":{"kvlistValue":{"values":[{"key":"Workstation Name","value":{"stringValue":"workstation"}},{"key":"Source Network Address","value":{"stringValue":"38.68.52.197"}},{"key":"Source Port","value":{"stringValue":"0"}}]}}},{"key":"Detailed Authentication Information","value":{"kvlistValue":{"values":[{"key":"Authentication Package","value":{"stringValue":"NTLM"}},{"key":"Transited Services","value":{"stringValue":"-"}},{"key":"Package Name (NTLM only)","value":{"stringValue":"-"}},{"key":"Key Length","value":{"stringValue":"0"}},{"key":"Logon Process","value":{"stringValue":"NtLmSsp"}}]}}},{"key":"Additional Context","value":{"arrayValue":{"values":[{"stringValue":"This event is generated when a logon request fails. It is generated on the computer where access was attempted."},{"stringValue":"The Subject fields indicate the account on the local system which requested the logon. This is most commonly a service such as the Server service, or a local process such as Winlogon.exe or Services.exe."},{"stringValue":"The Logon Type field indicates the kind of logon that was requested. The most common types are 2 (interactive) and 3 (network)."},{"stringValue":"The Process Information fields indicate which account and process on the system requested the logon."},{"stringValue":"The Network Information fields indicate where a remote logon request originated. Workstation name is not always available and may be left blank in some cases."},{"stringValue":"The authentication information fields provide detailed information about this specific logon request."},{"stringValue":"- Transited services indicate which intermediate services have participated in this logon request."},{"stringValue":"- Package name indicates which sub-protocol was used among the NTLM protocols."},{"stringValue":"- Key length indicates the length of the generated session key. This will be 0 if no session key was requested."}]}}},{"key":"Subject","value":{"kvlistValue":{"values":[{"key":"Security ID","value":{"stringValue":"S-1-0-0"}},{"key":"Account Name","value":{"stringValue":"-"}},{"key":"Account Domain","value":{"stringValue":"-"}},{"key":"Logon ID","value":{"stringValue":"0x0"}}]}}},{"key":"Logon Type","value":{"stringValue":"3"}},{"key":"Account For Which Logon Failed","value":{"kvlistValue":{"values":[{"key":"Security ID","value":{"stringValue":"S-1-0-0"}},{"key":"Account Name","value":{"stringValue":"Administrator"}},{"key":"Account Domain","value":{"stringValue":"-"}}]}}},{"key":"Failure Information","value":{"kvlistValue":{"values":[{"key":"Failure Reason","value":{"stringValue":"Unknown user name or bad password."}},{"key":"Status","value":{"stringValue":"0xC000006D"}},{"key":"Sub Status","value":{"stringValue":"0xC000006A"}}]}}},{"key":"Process Information","value":{"kvlistValue":{"values":[{"key":"Caller Process Name","value":{"stringValue":"-"}},{"key":"Caller Process ID","value":{"stringValue":"0x0"}}]}}}]}}},{"key":"event_id","value":{"kvlistValue":{"values":[{"key":"qualifiers","value":{"intValue":"0"}},{"key":"id","value":{"intValue":"4625"}}]}}},{"key":"channel","value":{"stringValue":"Security"}},{"key":"opcode","value":{"stringValue":"Info"}},{"key":"provider","value":{"kvlistValue":{"values":[{"key":"name","value":{"stringValue":"Microsoft-Windows-Security-Auditing"}},{"key":"guid","value":{"stringValue":"{54849625-5478-4994-a5ba-3e3b0328c30d}"}},{"key":"event_source","value":{"stringValue":""}}]}}}]}},"attributes":[{"key":"log_type","value":{"stringValue":"windows_event.security"}}],"traceId":"","spanId":""}]}]},{"resource":{"attributes":[{"key":"host.name","value":{"stringValue":"winhost-windows"}},{"key":"os.type","value":{"stringValue":"windows"}}]},"scopeLogs":[{"scope":{},"logRecords":[{"timeUnixNano":"1710783576180744200","observedTimeUnixNano":"1710775562753746900","severityNumber":9,"severityText":"INFO","body":{"kvlistValue":{"values":[{"key":"event_id","value":{"kvlistValue":{"values":[{"key":"qualifiers","value":{"intValue":"0"}},{"key":"id","value":{"intValue":"4625"}}]}}},{"key":"provider","value":{"kvlistValue":{"values":[{"key":"guid","value":{"stringValue":"{54849625-5478-4994-a5ba-3e3b0328c30d}"}},{"key":"event_source","value":{"stringValue":""}},{"key":"name","value":{"stringValue":"Microsoft-Windows-Security-Auditing"}}]}}},{"key":"channel","value":{"stringValue":"Security"}},{"key":"opcode","value":{"stringValue":"Info"}},{"key":"keywords","value":{"arrayValue":{"values":[{"stringValue":"Audit Failure"}]}}},{"key":"computer","value":{"stringValue":"winhost-windows"}},{"key":"message","value":{"stringValue":"An account failed to log on."}},{"key":"task","value":{"stringValue":"Logon"}},{"key":"record_id","value":{"intValue":"107867"}},{"key":"details","value":{"kvlistValue":{"values":[{"key":"Subject","value":{"kvlistValue":{"values":[{"key":"Security ID","value":{"stringValue":"S-1-0-0"}},{"key":"Account Name","value":{"stringValue":"-"}},{"key":"Account Domain","value":{"stringValue":"-"}},{"key":"Logon ID","value":{"stringValue":"0x0"}}]}}},{"key":"Logon Type","value":{"stringValue":"3"}},{"key":"Account For Which Logon Failed","value":{"kvlistValue":{"values":[{"key":"Security ID","value":{"stringValue":"S-1-0-0"}},{"key":"Account Name","value":{"stringValue":"ADMINISTRATOR"}},{"key":"Account Domain","value":{"stringValue":"-"}}]}}},{"key":"Failure Information","value":{"kvlistValue":{"values":[{"key":"Failure Reason","value":{"stringValue":"Unknown user name or bad password."}},{"key":"Status","value":{"stringValue":"0xC000006D"}},{"key":"Sub Status","value":{"stringValue":"0xC000006A"}}]}}},{"key":"Process Information","value":{"kvlistValue":{"values":[{"key":"Caller Process ID","value":{"stringValue":"0x0"}},{"key":"Caller Process Name","value":{"stringValue":"-"}}]}}},{"key":"Network Information","value":{"kvlistValue":{"values":[{"key":"Workstation Name","value":{"stringValue":"-"}},{"key":"Source Network Address","value":{"stringValue":"115.243.85.101"}},{"key":"Source Port","value":{"stringValue":"0"}}]}}},{"key":"Detailed Authentication Information","value":{"kvlistValue":{"values":[{"key":"Logon Process","value":{"stringValue":"NtLmSsp"}},{"key":"Authentication Package","value":{"stringValue":"NTLM"}},{"key":"Transited Services","value":{"stringValue":"-"}},{"key":"Package Name (NTLM only)","value":{"stringValue":"-"}},{"key":"Key Length","value":{"stringValue":"0"}}]}}},{"key":"Additional Context","value":{"arrayValue":{"values":[{"stringValue":"This event is generated when a logon request fails. It is generated on the computer where access was attempted."},{"stringValue":"The Subject fields indicate the account on the local system which requested the logon. This is most commonly a service such as the Server service, or a local process such as Winlogon.exe or Services.exe."},{"stringValue":"The Logon Type field indicates the kind of logon that was requested. The most common types are 2 (interactive) and 3 (network)."},{"stringValue":"The Process Information fields indicate which account and process on the system requested the logon."},{"stringValue":"The Network Information fields indicate where a remote logon request originated. Workstation name is not always available and may be left blank in some cases."},{"stringValue":"The authentication information fields provide detailed information about this specific logon request."},{"stringValue":"- Transited services indicate which intermediate services have participated in this logon request."},{"stringValue":"- Package name indicates which sub-protocol was used among the NTLM protocols."},{"stringValue":"- Key length indicates the length of the generated session key. This will be 0 if no session key was requested."}]}}}]}}},{"key":"execution","value":{"kvlistValue":{"values":[{"key":"process_id","value":{"intValue":"660"}},{"key":"thread_id","value":{"intValue":"3244"}}]}}},{"key":"system_time","value":{"stringValue":"2024-03-18T15:26:00.9126673Z"}},{"key":"level","value":{"stringValue":"Information"}},{"key":"event_data","value":{"kvlistValue":{"values":[{"key":"data","value":{"arrayValue":{"values":[{"kvlistValue":{"values":[{"key":"SubjectUserSid","value":{"stringValue":"S-1-0-0"}}]}},{"kvlistValue":{"values":[{"key":"SubjectUserName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"SubjectDomainName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"SubjectLogonId","value":{"stringValue":"0x0"}}]}},{"kvlistValue":{"values":[{"key":"TargetUserSid","value":{"stringValue":"S-1-0-0"}}]}},{"kvlistValue":{"values":[{"key":"TargetUserName","value":{"stringValue":"ADMINISTRATOR"}}]}},{"kvlistValue":{"values":[{"key":"TargetDomainName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"Status","value":{"stringValue":"0xc000006d"}}]}},{"kvlistValue":{"values":[{"key":"FailureReason","value":{"stringValue":"%%2313"}}]}},{"kvlistValue":{"values":[{"key":"SubStatus","value":{"stringValue":"0xc000006a"}}]}},{"kvlistValue":{"values":[{"key":"LogonType","value":{"stringValue":"3"}}]}},{"kvlistValue":{"values":[{"key":"LogonProcessName","value":{"stringValue":"NtLmSsp "}}]}},{"kvlistValue":{"values":[{"key":"AuthenticationPackageName","value":{"stringValue":"NTLM"}}]}},{"kvlistValue":{"values":[{"key":"WorkstationName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"TransmittedServices","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"LmPackageName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"KeyLength","value":{"stringValue":"0"}}]}},{"kvlistValue":{"values":[{"key":"ProcessId","value":{"stringValue":"0x0"}}]}},{"kvlistValue":{"values":[{"key":"ProcessName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"IpAddress","value":{"stringValue":"115.243.85.101"}}]}},{"kvlistValue":{"values":[{"key":"IpPort","value":{"stringValue":"0"}}]}}]}}}]}}}]}},"attributes":[{"key":"log_type","value":{"stringValue":"windows_event.security"}}],"traceId":"","spanId":""}]}]},{"resource":{"attributes":[{"key":"host.name","value":{"stringValue":"winhost-windows"}},{"key":"os.type","value":{"stringValue":"windows"}}]},"scopeLogs":[{"scope":{},"logRecords":[{"timeUnixNano":"1710783587483696000","observedTimeUnixNano":"1710775573749320300","severityNumber":9,"severityText":"INFO","body":{"kvlistValue":{"values":[{"key":"execution","value":{"kvlistValue":{"values":[{"key":"process_id","value":{"intValue":"660"}},{"key":"thread_id","value":{"intValue":"3244"}}]}}},{"key":"channel","value":{"stringValue":"Security"}},{"key":"record_id","value":{"intValue":"107868"}},{"key":"level","value":{"stringValue":"Information"}},{"key":"message","value":{"stringValue":"An account failed to log on."}},{"key":"task","value":{"stringValue":"Logon"}},{"key":"keywords","value":{"arrayValue":{"values":[{"stringValue":"Audit Failure"}]}}},{"key":"computer","value":{"stringValue":"winhost-windows"}},{"key":"event_data","value":{"kvlistValue":{"values":[{"key":"data","value":{"arrayValue":{"values":[{"kvlistValue":{"values":[{"key":"SubjectUserSid","value":{"stringValue":"S-1-0-0"}}]}},{"kvlistValue":{"values":[{"key":"SubjectUserName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"SubjectDomainName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"SubjectLogonId","value":{"stringValue":"0x0"}}]}},{"kvlistValue":{"values":[{"key":"TargetUserSid","value":{"stringValue":"S-1-0-0"}}]}},{"kvlistValue":{"values":[{"key":"TargetUserName","value":{"stringValue":"ADMINISTRATOR"}}]}},{"kvlistValue":{"values":[{"key":"TargetDomainName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"Status","value":{"stringValue":"0xc000006d"}}]}},{"kvlistValue":{"values":[{"key":"FailureReason","value":{"stringValue":"%%2313"}}]}},{"kvlistValue":{"values":[{"key":"SubStatus","value":{"stringValue":"0xc000006a"}}]}},{"kvlistValue":{"values":[{"key":"LogonType","value":{"stringValue":"3"}}]}},{"kvlistValue":{"values":[{"key":"LogonProcessName","value":{"stringValue":"NtLmSsp "}}]}},{"kvlistValue":{"values":[{"key":"AuthenticationPackageName","value":{"stringValue":"NTLM"}}]}},{"kvlistValue":{"values":[{"key":"WorkstationName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"TransmittedServices","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"LmPackageName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"KeyLength","value":{"stringValue":"0"}}]}},{"kvlistValue":{"values":[{"key":"ProcessId","value":{"stringValue":"0x0"}}]}},{"kvlistValue":{"values":[{"key":"ProcessName","value":{"stringValue":"-"}}]}},{"kvlistValue":{"values":[{"key":"IpAddress","value":{"stringValue":"207.249.123.189"}}]}},{"kvlistValue":{"values":[{"key":"IpPort","value":{"stringValue":"0"}}]}}]}}}]}}},{"key":"system_time","value":{"stringValue":"2024-03-18T15:26:12.2156191Z"}},{"key":"details","value":{"kvlistValue":{"values":[{"key":"Account For Which Logon Failed","value":{"kvlistValue":{"values":[{"key":"Security ID","value":{"stringValue":"S-1-0-0"}},{"key":"Account Name","value":{"stringValue":"ADMINISTRATOR"}},{"key":"Account Domain","value":{"stringValue":"-"}}]}}},{"key":"Failure Information","value":{"kvlistValue":{"values":[{"key":"Failure Reason","value":{"stringValue":"Unknown user name or bad password."}},{"key":"Status","value":{"stringValue":"0xC000006D"}},{"key":"Sub Status","value":{"stringValue":"0xC000006A"}}]}}},{"key":"Process Information","value":{"kvlistValue":{"values":[{"key":"Caller Process ID","value":{"stringValue":"0x0"}},{"key":"Caller Process Name","value":{"stringValue":"-"}}]}}},{"key":"Network Information","value":{"kvlistValue":{"values":[{"key":"Workstation Name","value":{"stringValue":"-"}},{"key":"Source Network Address","value":{"stringValue":"207.249.123.189"}},{"key":"Source Port","value":{"stringValue":"0"}}]}}},{"key":"Detailed Authentication Information","value":{"kvlistValue":{"values":[{"key":"Logon Process","value":{"stringValue":"NtLmSsp"}},{"key":"Authentication Package","value":{"stringValue":"NTLM"}},{"key":"Transited Services","value":{"stringValue":"-"}},{"key":"Package Name (NTLM only)","value":{"stringValue":"-"}},{"key":"Key Length","value":{"stringValue":"0"}}]}}},{"key":"Additional Context","value":{"arrayValue":{"values":[{"stringValue":"This event is generated when a logon request fails. It is generated on the computer where access was attempted."},{"stringValue":"The Subject fields indicate the account on the local system which requested the logon. This is most commonly a service such as the Server service, or a local process such as Winlogon.exe or Services.exe."},{"stringValue":"The Logon Type field indicates the kind of logon that was requested. The most common types are 2 (interactive) and 3 (network)."},{"stringValue":"The Process Information fields indicate which account and process on the system requested the logon."},{"stringValue":"The Network Information fields indicate where a remote logon request originated. Workstation name is not always available and may be left blank in some cases."},{"stringValue":"The authentication information fields provide detailed information about this specific logon request."},{"stringValue":"- Transited services indicate which intermediate services have participated in this logon request."},{"stringValue":"- Package name indicates which sub-protocol was used among the NTLM protocols."},{"stringValue":"- Key length indicates the length of the generated session key. This will be 0 if no session key was requested."}]}}},{"key":"Subject","value":{"kvlistValue":{"values":[{"key":"Security ID","value":{"stringValue":"S-1-0-0"}},{"key":"Account Name","value":{"stringValue":"-"}},{"key":"Account Domain","value":{"stringValue":"-"}},{"key":"Logon ID","value":{"stringValue":"0x0"}}]}}},{"key":"Logon Type","value":{"stringValue":"3"}}]}}},{"key":"event_id","value":{"kvlistValue":{"values":[{"key":"qualifiers","value":{"intValue":"0"}},{"key":"id","value":{"intValue":"4625"}}]}}},{"key":"provider","value":{"kvlistValue":{"values":[{"key":"name","value":{"stringValue":"Microsoft-Windows-Security-Auditing"}},{"key":"guid","value":{"stringValue":"{54849625-5478-4994-a5ba-3e3b0328c30d}"}},{"key":"event_source","value":{"stringValue":""}}]}}},{"key":"opcode","value":{"stringValue":"Info"}}]}},"attributes":[{"key":"log_type","value":{"stringValue":"windows_event.security"}}],"traceId":"","spanId":""}]}]}]}`,
	},
}