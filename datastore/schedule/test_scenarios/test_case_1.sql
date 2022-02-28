INSERT INTO schedules (id, slack_channel_id, team_slack_id, name, interval, is_enabled, end_time, start_time, weekdays_only, created_at, updated_at) VALUES
(1,'C02FZ54HJP6','T022YEG9FGA','Daily On Call','daily',true,'16:00:00+00:00','08:00:00+00:00',true,'2022-02-20 20:23:26.200378+00','2022-02-20 20:23:26.200378+00'),
(2,'XASLKASDJ8','T0YUBV18290','On Call Schedule','weekly',true,'16:00:00+00:00','08:00:00+00:00',false,'2022-02-20 20:23:26.200378+00','2022-02-20 20:23:26.200378+00');

INSERT INTO users (id,slack_id,slack_handle,email,first_name,last_name,avatar_url,display_name,created_at,updated_at) VALUES
(1,'U030NTQ1G73','<@U030NTQ1G73|jon.smith>','jon.smith@test.io','Jon','Smith','https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png','Jon Smith','2022-02-20 20:23:46.175331+00','2022-02-20 20:23:46.175331+00'),
(2,'U02QRKM83GA','<@U02QRKM83GA|mary.black>','mary.black@test.io','Mary','Black','https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png','Mary Black','2022-02-20 20:23:54.05764+00','2022-02-20 20:23:54.05764+00'),
(3,'U02SGMWHN75','<@U02SGMWHN75|jenny.bleeker>','jenny.bleeker@test.io','Jenny','Bleeker','https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png','Jenny Bleeker','2022-02-20 20:24:33.935131+00','2022-02-20 20:24:33.935131+00'),
(4,'U02PJ4ZNE8H','<@U02PJ4ZNE8H|jane.doe>','jane.doe@test.io','Jane','Doe','https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png','Jane Doe','2022-02-20 20:24:44.383539+00','2022-02-20 20:24:44.383539+00'),
(5,'U02QRLQU09H','<@U02QRLQU09H|jack.renshaw>','jack.renshaw@test.io','Jack','Renshaw','https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png','Jack Renshaw','2022-02-20 20:24:55.032375+00','2022-02-20 20:24:55.032375+00'),
(6,'U0303JLF1QW','<@U0303JLF1QW|jon.doe>','jon.doe@test.io','Jon','Doe','https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png','Jon Doe','2022-02-20 20:25:12.783014+00','2022-02-20 20:25:12.783014+00');

INSERT INTO shifts (sequence_id,user_id,schedule_id,status,started_at,created_at,updated_at) VALUES
(8,6,1,null,null,'2022-02-20 20:25:12.789541+00','2022-02-20 20:25:12.789541+00'),
(12,2,1,null,null,'2022-02-20 20:23:54.064455+00','2022-02-23 08:01:00.034545+00'),
(23,3,1,null,null,'2022-02-20 20:24:33.941838+00','2022-02-24 08:01:00.011914+00'),
(6,4,1,null,null,'2022-02-20 20:24:44.390259+00','2022-02-25 08:01:00.007043+00'),
(5,5,1,'active','2022-02-25 08:01:00.0065+00','2022-02-20 20:24:55.039688+00','2022-02-20 20:24:55.039688+00'),
(1,1,1,'override','2022-02-25 09:26:22.788711+00','2022-02-20 20:23:46.183879+00','2022-02-25 09:26:22.788994+00');